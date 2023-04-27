cd ~/Developer/personal-website

echo "Updating Caddy server with new config 🛠"
echo "  Uploading Caddyfile to server..."
scp Caddyfile root@blog.michaelmiranda.dev:/etc/caddy/michaelmiranda.dev
echo "  Caddyfile uploaded sucessfully 📝"

echo "  Reloading Caddy server..."
ssh root@blog.michaelmiranda.dev "systemctl reload caddy"
echo "  Caddy reloaded successfully 🔁"
echo "Caddy server updated 🎉"

