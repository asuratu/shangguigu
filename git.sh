CONTENT=$1

# CONTENT 默认值为 "update"
if [ -z "$CONTENT" ]; then
    CONTENT="update"
fi

git add .
git commit -m "$CONTENT"
git pull
git push

printf "\033[32m 代码提交成功! \033[0m\n"
