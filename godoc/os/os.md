# type

struct fileStat (file info)  
interface FileInfo  
FileMode  

# sys  
hostname() /proc/sys/kernel/hosthome  
(like z02)

# str
itoa 

# stat
file info compare or get  

# proc
get uid euid gid egid group  

# path
mkdir and removedir   
 
# getwd
return current dir 

# file
file operation(read,write,name)  
file {fd, name, dirinfo, nepipe}  
Truncate(name,size)截取文件中的一段文字  
basename(name) 去掉前面的文件夹，保留文件名  
Pipe() 

# env
Getenv(str) get env info var name  
unset  
set  
