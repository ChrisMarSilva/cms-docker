# sed -i  "5a ccc" file
# "5a"表示在第5行之后添加，"ccc"表示要添加的内容，若想在第5行之前添加内容可以把"a"换成"i"
sed -i "5^s #" /etc/mysql/my.cnf
# sed -i "7^ #" /etc/mysql/my.cnf
# sed -i "11^ #" /etc/mysql/my.cnf
cat /etc/mysql/my.cnf
