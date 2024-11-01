echo "Fetching branch"
git pull origin main

echo "Redeploying"
./redeploy.sh