import time
import subprocess
import shlex
import json
from sqlihter import SQLighter
import schedule
from threading import Thread
import platform

shell_status = True
system_check = platform.system()
working_directory = 'rclone'

def run_command(command):
    json_data = ''

    process = subprocess.Popen(shlex.split(command), shell=shell_status, stdout=subprocess.PIPE, cwd=working_directory)
    while True:
        output = process.stdout.readline().decode('utf-8')
        if output == '' and process.poll() is not None:
            break
        if output:
            # print (output)
            # print ('output: ', output.strip())
            json_data += output
        process.poll()
    # print ('rc:', rc)
    # print (json_data)
    return json_data
    # return rc

if system_check == 'Linux':

    shell_status = False
    working_directory = 'rclone-v1.57.0-linux-amd64'
    run_command('chmod +x rclone')
    run_command('./rclone config --config="rclone.conf"')


def checker(db, cloud):
    t0 = time.time()
    print(cloud)

    json_txt = run_command(f'./rclone lsjson {cloud}: --recursive')

    # db = SQLighter("dc25luei5qje0r")

    # print('json_txt: ', json_txt)

    t1 = time.time()
    print("Time elapsed: ", t1 - t0)

    json_new = json.loads(json_txt)

    json_cloud = db.send_json(cloud)
    # print (len(json_cloud))
    if len(json_cloud) == 0:
        db.add_cloud(cloud, json_txt)
        return cloud
    # print ('json_cloud: ', json_cloud)
    json_old = json.loads(json_cloud)

    ids = []
    modtime = []
    names = []
    path = []
    ids_path_new = []
    for new in json_new:
        ids.append(new['ID'])
        modtime.append(new['ModTime'])
        names.append(new['Name'])
        path.append(new['Path'])

        ids_path = []
        path_list = new['Path'].split('/')
        # print(path_list)
        # print(len(path_list))
        if len(path_list) > 1:
            for i in path_list:
                for new_check in json_new:
                    if new_check['Name'] == i:
                        ids_path.append(new['ID'])
            strA = "/".join(ids_path)
            ids_path_new.append(strA)
            # print(strA)
    # print ("ids_path_new: ", ids_path_new)

    count = 0
    for old in json_old:
        statuses = []

        if old['ID'] not in ids:
            # не будет работать в ситуациях, если файл будет восстановлен, а потом опять удален между проверками
            count += 1
            status = "delete"
            statuses.append(status)

        if old['ModTime'] not in modtime and old['ID'] in ids and old['Name'] in names:
            if cloud == "dropbox_test" and old["MimeType"] == "inode/directory":
                # print ('Вышел на папку')
                pass
            else:
                count += 1
                status = "modify"
                statuses.append(status)

        if old['Name'] not in names and old['ID'] in ids:
            count += 1
            status = "rename"
            statuses.append(status)

        if old['Path'] not in path and old['ID'] in ids and old['Name'] in names:
            ids_path_old = []
            path_list = old['Path'].split('/')
            # Внимание! Жуткий костыль
            if len(path_list) > 1:
                for i in path_list:
                    for old_check in json_old:
                        if old_check['Name'] == i:
                            ids_path_old.append(old['ID'])
                id_path_old = "/".join(ids_path_old)
                # print (id_path_old)
                if id_path_old not in ids_path_new:
                    continue
                else:
                    break
            status = "move"
            statuses.append(status)
            # не будет работать в ситуациях, когда файл между проверками будет перемещен, и переименован (или наоборот)

        for status in statuses:
            print(f"Файл был удален под названием {old['Name']} был {status}. \nID:{old['ID']}\n\n")
            json_data = [{f"{cloud}": {"file_id": old['ID'], "filename": old['Path'], "time_changed": old['ModTime'],
                                       "status": {status}}}]
            json_data = json.dumps(json_data, sort_keys=True, indent=2)
            # print(type(json_data))
            print(json_data)
            db.add_json(json_data)

    if count == 0:
        print('Изменений нет')

    db.delete_table(cloud)
    db.add_cloud(cloud, json_txt)
    db.close()


def clouds():
    clouds = ["google_drive_test", "one_drive_test", "dropbox_test"]
    for cloud in clouds:
        checker(cloud)
        db = SQLighter("dc25luei5qje0r")
        th = Thread(target=checker, args=(db, cloud,))
        th.start()

    # print ("следующая очередь")


schedule.every(30).seconds.do(clouds)

while True:
    schedule.run_pending()
    time.sleep(1)
