�����������

������ƣ������񡤹����Զ�����ϵͳ

������ܣ��Զ�����mysql���ݿ��е�<script></script>������룡




���ֹ�������

�������./hwsmysqlclear run -u root -p ���� -db ���ݿ��� -t 30

ע�⣺ÿ����������������Ҫ�ֹ����������




������������

1���޸�hwsmysqlcleard�е����ݿ�������Ϣ

2��ע�Ტ���������������sudo make install


ע�⣺ÿ����������������Զ����������

ж�ط�����sudo make uninstall




������˵����

������������鿴������./hwsmysqlclear help run

Usage of command "run":

        hwsmysqlclear run [options]

Options:

  -db string
        database name
  -exclude string
        Exclude tables, comma separated
  -include string
        Include tables, comma separated
  -p string
        password
  -t int
        How many seconds between scans (default 10)
  -u string
        username (default "root")


