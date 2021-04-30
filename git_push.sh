#!/bin/bash
echo "初始化中....."
git pull origin main;
git add -A;
read -p "输入日志,按Enter键跳过:" log
if  [ ! -n "$log" ] ;then
    git commit -m "自动生成";
else git commit -m "${log}";
fi;
git branch -M main
git remote add origin git@github.com:cyylog/cyylgo.git
git push -u origin main
echo "远程推送完成"



