import art
import hep_test
from os import system, name
from time import sleep


def clear():
    # for windows
    if name == 'nt':
        _ = system('cls')
    # for mac and linux(here, os.name is 'posix')
    else:
        _ = system('clear')

dict = [
    {'num': 1, 'test': 'Anti-Bot', 'function': hep_test.antibot_test},
    {'num': 2, 'test': 'web download Anti-Virus', 'function': hep_test.avtest},
    {'num': 3, 'test': 'Zero-Phishing', 'function': hep_test.zero_phishing},
    {'num': 4, 'test': 'on-access Anti-virus Scanning', 'function': hep_test.on_access_av},
    {'num': 5, 'test': 'on-access Anti-virus Scanning - 2', 'function': hep_test.on_access_av_base64},
    {'num': 6, 'test': 'Local Threat Emulation', 'function': hep_test.revshell_te},
    {'num': 7, 'test': 'Behaviour Guard', 'function': hep_test.cred_dump},
    {'num': 8, 'test': 'URL filtering', 'function': hep_test.url_filtering},
    {'num': 9, 'test': 'Threat Hunting', 'function': hep_test.sch_task}
    ]


if __name__ == '__main__':
    end = False
    while not end:
        clear()
        print(art.logo)
        print("Welcome to the Harmony Testing Tool!")
        print("Test Case: Description")
        for i in range(0, len(dict)):
            print(str(dict[i]['num']) + ": " + dict[i]['test'])
        # testnum = int(input("Please select the test case number for starting the respective test: ")) - 1
        correctnum = False
        while not correctnum:
            testnum = int(input("Please select the test case number for starting the respective test: ")) - 1
            # print(testnum)
            if testnum in (range(0,len(dict))):
                # print("ok")
                dict[testnum]['function']()
                correctnum = True
            else:
                print("Please re-enter the test case number! ")
                correctnum = False
        ending = input("Type 'cont' to continue testing other cases, or type 'exit' to quit: ")
        if ending != 'cont':    
            end = True
