---
date: "2025-07-04"
title: "Building a lightweight React hook for reCAPTCHA integration"
description: "Creating a zero dependency reCAPTCHA hook to limit bots in your react application."
---

#### _Disclaimer_

_Technically there are dependencies as this depends on React and reCAPTCHA but you get the gist of it &#128517;_

### What is reCAPTCHA?

[reCAPTCHA](https://www.google.com/recaptcha/about/) is a free service from Google that helps protect your website from spam and abuse. By integrating reCAPTCHA into your React application, you can help ensure that only humans are able to interact with your forms and other sensitive content.

If you find yourself needing to add reCAPTCHA to your application you might think about reaching for yet another library to handle that. However, you can create a reusable hook for your application in a few simple steps.

### Prerequisites

To follow along, you should have:

- A decent understanding of React (mainly hooks)
- Created a pair of [reCAPTCHA keys](https://www.google.com/recaptcha/admin/create)
  - _Note: in this article we'll be using reCAPTCHA v3_
- A server side application that can process the generated token
- TypeScript (optional)

### Let's get started

In this example I'll be using TypeScript so the first thing we want to do is make our compiler happy! To avoid being yelled at, let's create an `index.d.ts` file with the following:

```ts
declare global {
  interface Window {
    grecaptcha?: {
      execute: (widgetId?: number) => Promise<string>;
      getResponse: (widgetId?: number) => string;
      render: (
        container: string | HTMLElement,
        params: {
          badge?: "bottomright" | "bottomleft" | "inline";
          callback?: (recaptchaResponse: string) => void;
          "error-callback"?: (recaptchaResponse: string) => void;
          "expired-callback"?: (recaptchaResponse: string) => void;
          isolated?: boolean;
          sitekey: string;
          size?: "normal" | "compact" | "invisible";
          theme?: "dark" | "light";
        }
      ) => number;
      reset: (widgetId?: number) => void;
    };
  }
}
```

The type definition we just created is based on the reCAPTCHA v2 [documentation](https://developers.google.com/recaptcha/docs/invisible). While the documentation is for v2, it is still applicable to v3.

You may have noticed that the type definition does not exactly match the documentation. This is because the documentation is not up to date. However, I have tested this type definition and it works well. If you are curious, you can test it yourself.

Now that TypeScript is happy, let's define the API for our hook. To support the most important functionality reCAPTCHA provides, we can create the following type.

```ts
type UseRecaptchaProps = {
  badge?: "bottomright" | "bottomleft" | "inline";
  containerId: string;
  onError?: (recaptchaResponse: string) => void;
  onExpired?: (recaptchaResponse: string) => void;
  onSuccess?: (recaptchaResponse: string) => void;
  shouldLoad?: boolean;
  theme?: "dark" | "light";
};
```

To use this hook properly, we only need to pass it a `containerId`. This is the id of the DOM element that will serve as our reCAPTCHA badge. We can provide defaults for the rest of the props or not use them at all.

You'll also notice that I've included a `shouldLoad` prop. This prop is not documented by reCAPTCHA, but I added it to give us the flexibility of determining when to load reCAPTCHA. This way, we can save bandwidth for users who may never interact with our forms.

### Creating our hook

```tsx
const useRecaptcha = ({
  badge,
  containerId,
  onError,
  onExpired,
  onSuccess,
  shouldLoad = true, // allows us to delay loading of recaptcha (lets say until a user interacts with a form)
  theme = "light"
}: UseRecaptchaProps) => {
  const isClient = typeof window !== "undefined";
  const [recaptchaLoaded, setRecaptchaLoaded] = useState(
    isClient && window?.grecaptcha?.render ? true : false
  );
  const [recaptchaWidget, setRecaptchaWidget] = useState<number | null>(null);
  const checkLoadRef = useRef<number>();

  useEffect(() => {
    if (shouldLoad) {
      const script = document.createElement("script");
      script.src = "https://www.google.com/recaptcha/api.js?render=explicit";
      script.async = true;
      script.defer = true;
      document.body.appendChild(script);

      checkLoadRef.current = window.setInterval(() => {
        if (window?.grecaptcha?.render) {
          setRecaptchaLoaded(true);
        }
      }, 300);
    }

    return () => {
      clearInterval(checkLoadRef.current);
    };
  }, [shouldLoad]);

  useEffect(() => {
    if (recaptchaLoaded && window.grecaptcha && recaptchaWidget === null) {
      clearInterval(checkLoadRef.current);
      const widget = window.grecaptcha.render(containerId, {
        badge,
        callback: onSuccess,
        "error-callback": onError,
        "expired-callback": onExpired,
        sitekey: GOOGLE_RECAPTCHA_SITE_KEY,
        size: "invisible",
        theme
      });

      setRecaptchaWidget(widget);
    }
  }, [recaptchaLoaded, onSuccess, recaptchaWidget, containerId]); // eslint-disable-line react-hooks/exhaustive-deps

  const executeRecaptcha = useCallback(async () => {
    if (recaptchaWidget !== null && window.grecaptcha) {
      // Do not assume a successful form submission! Always reset before executing.
      window.grecaptcha.reset(recaptchaWidget);
      const token = await window.grecaptcha.execute(recaptchaWidget);
      return token;
    }
  }, [recaptchaWidget]);

  return {
    executeRecaptcha,
    recaptchaLoaded
  };
};
```

One of the most important features of this hook is that it allows us to asynchronously call the `grecaptcha.execute()` method. This is not obvious from the documentation, but it is a very powerful feature.

By calling `grecaptcha.execute()` asynchronously, we can avoid having to wrap any existing business logic inside of the reCAPTCHA callback (`onSuccess`). Instead, we can call our own `executeRecaptcha()` function from anywhere in our code. This gives us much more flexibility and control over how we interact with reCAPTCHA.

Here is an example of how we can use our hook. We will also be using the `shouldLoad` property that I mentioned earlier to improve the performance of our application.

```tsx
import { useState } from "react";
import { useRecaptcha } from "../hooks";

const Form = () => {
  const recaptchaContainerId = "recaptcha-container";
  const [enableRecaptcha, setEnableRecaptcha] = useState(false);
  const { executeRecaptcha, recaptchaLoaded } = useRecaptcha({
    containerId: recaptchaContainerId,
    shouldLoad: enableRecaptcha
  });

  return (
    <form
      onSubmit={async (e) => {
        e.preventDefault();
        if (enableRecaptcha) {
          const token = await executeRecaptcha();
          // You'd want to pass this token to your backend which
          // isn't covered in this article so we'll just log it for now.
          console.log("recaptchaToken:", token);
        }
      }}
    >
      {enableRecaptcha && (
        <div
          id={recaptchaContainerId}
          style={{ position: "absolute", zIndex: "9999" }}
        />
      )}
      <label for="email">Email:</label>
      <input
        id="email"
        name="email"
        onFocus={() => {
          setEnableRecaptcha(true);
        }}
        type="text"
      />
      <button disabled={recaptchaLoaded} type="submit">
        Submit
      </button>
    </form>
  );
};
```

Again, using the `shouldLoad` property is optional, so we can simplify this code even further by removing that logic.

As a prerequisite, I mentioned that you should already have an API in place that will process the token we just generated. If you don't have one in place you can take a look at the [documentation](https://developers.google.com/recaptcha/docs/verify) to properly validate the user's token on the server.

Well there you have it! We've successfully created a reusable hook that can be used to generate a reCAPTCHA token throughout our application without relying on a third-party library.
