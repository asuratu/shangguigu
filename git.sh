CONTENT=$1

# CONTENT 默认值为 "update"
if [ -z "$CONTENT" ]; then
    CONTENT="update"
fi

git add .
git commit -m "$CONTENT"

git pull
echo "git pull success"

git push origin main
echo "git push success"
exit
