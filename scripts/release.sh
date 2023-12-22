#!/bin/bash
cd ~/Developer/personal-website
rm -rf build
mkdir build

echo "Releasing personal-website 🏁"
echo "  Building the server..."
GOOS=linux GOARCH=amd64 go build -o ./build/server .
echo "  Server built successfully 🏗"

echo "  Bundling code..."
cp -R database/migrations public .env.production Caddyfile build/
echo "  Bundle created successfully 📦"

echo "  Minifying CSS..."
./bin/tailwindcss -m -i tailwind.css -o build/public/styles/site.css
echo " CSS Minified Successfully 🎨"

echo "  Deleting existing code on server..."
ssh root@michaelmiranda.dev "rm -rf /home/mike/apps/personal-website"
echo "  Code deleted successfully ✅"

echo "  Uploading code..."
rsync -avr build/ root@michaelmiranda.dev:/home/mike/apps/personal-website
echo "  Code uploaded successfully 📦"

echo "  Cleaning up..."
rm -rf build
echo "  Finished cleaning 🧹"

echo "  Moving Caddyfile..."
ssh root@michaelmiranda.dev "cp /home/mike/apps/personal-website/Caddyfile /etc/caddy/michaelmiranda.dev"
echo "  Caddyfile moved successfully 🚚"

echo "  Restarting Application server..."
ssh root@michaelmiranda.dev "service michaelmiranda.dev restart"
echo "  Server restarted successfully 🔁"

echo "  Reloading Caddy server..."
ssh root@michaelmiranda.dev "systemctl reload caddy"
echo "  Caddy reloaded successfully 🔁"
echo "Release finished 🚀"