#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
reload(sys)
from email.MIMEText import MIMEText
import smtplib
sys.setdefaultencoding('utf-8')
import socket, fcntl, struct

def send_mail(to_list,sub,content):
 mail_host="smtp.126.com"    #使用的邮箱的smtp服务器地址
 mail_user="wudebao5220150@126.com"   #用户名
 mail_pass="wdb5221461121319"   #密码
 mail_postfix="126.com"   #邮箱的后缀
 me=mail_user+"<"+mail_user+"@"+mail_postfix+">"
 msg = MIMEText(content)
 msg['Subject'] = sub
 msg['From'] = me
 msg['To'] = to_list
 try:
  s = smtplib.SMTP()
  s.connect(mail_host)
  s.login(mail_user,mail_pass)
  s.sendmail(me, to_list, msg.as_string())
  s.close()
  return True
 except Exception, e:
  print str(e)
  return False
def get_local_ip(ifname = 'eth1'):
 s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
 inet = fcntl.ioctl(s.fileno(), 0x8915, struct.pack('256s', ifname[:15]))
 ret = socket.inet_ntoa(inet[20:24])
 return ret
if sys.argv[1]!="master" and sys.argv[1]!="backup" and sys.argv[1]!="fault":  #获取参数
 sys.exit()
else:
 notify_type = sys.argv[1]

if __name__ == '__main__':
 strcontent = get_local_ip()+ " " +notify_type+" State is activated, please make sure the HAProxy service running state!"
 mailto_list = ['472119740@qq.com']     #收件人(列表)
for mailto in mailto_list:
 send_mail(mailto, "HAProxy State switch alarm", strcontent.encode('utf-8'))


