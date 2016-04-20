# bapi

使用beego开发的api,实现从目录获取文件中的jason格式的data，并进行简单分析后进行回吐
<br>提供POST、GET、GET[reportid]、DELETE功能
<br>1、获取代码
<br>2、进入bapi目录执行“bee run -downdoc=true -gendoc=true”命令
<br>3、打开页面http://127.0.0.1:8080/swagger/swagger-1/

POST 请求body格式
{
"ReportId:" "YYYYMM"
}
<br>GET 请求格式
reportid "YYYYMM"
<br>DELETE 请求格式
reportid "YYYYMM"
