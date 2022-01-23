import sys
from simple_image_download import simple_image_download

# args[1] = 검색어 args[2] = 갯수
args = sys.argv
response = simple_image_download.simple_image_download()

if len(args) != 3:
    raise Exception('파라미터 갯수를 확인해라')

keyword = args[1]
limit = int(args[2])

if limit < 1 or limit > 10:
    raise Exception('갯수를 작작 틀려야지')

urls = response.urls(keyword, limit+2)

for i in urls:
    if not(i.__contains__('https://www.gstatic.com/ui/v1/menu/dark_thumbnail2.png')  or i.__contains__('https://www.gstatic.com/ui/v1/menu/device_default_thumbnail2.png')):
        print(i)