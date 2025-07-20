<script>
  const birthDate = new Date(1996, 11, 19);

  const calculateAge = () => {
    const today = new Date();
    let age = today.getFullYear() - birthDate.getFullYear();
    const monthDiff = today.getMonth() - birthDate.getMonth();

    // If birthday hasn't occurred yet this year, subtract 1
    if (
      monthDiff < 0 ||
      (monthDiff === 0 && today.getDate() < birthDate.getDate())
    ) {
      age--;
    }

    return age;
  };

  const age = $derived(calculateAge());
</script>

<div class="mx-auto min-h-screen max-w-4xl md:mt-44">
  <h1
    class="mb-10 text-4xl font-medium leading-tight text-slate-900 md:text-5xl dark:text-white"
  >
    About Michael
  </h1>
  <p class="text-lg md:text-xl">
    Hey! My name is Michael Miranda. I am a {age} year old software engineer based
    in NYC 🗽
  </p>
</div>
