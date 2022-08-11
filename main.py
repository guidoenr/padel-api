import requests
import colors
from bs4 import BeautifulSoup
from datetime import datetime


def parse_button(button_title):
    splitted = button_title.split()
    date = splitted[5]
    hour = splitted[8]

    date_time_obj = datetime.strptime(date, '%d/%m/%Y')
    day_name = date_time_obj.strftime('%A')
    
    return day_name, date, hour

def get_turnos(field):
    page = requests.get(field)
    soup = BeautifulSoup(page.content, "html.parser")
    buttons = soup.find_all("button") # buttons are the Tags that contains all the info about a turno
    initial_day, date, hour = parse_button(b["title"])
    for b in buttons:
        colors.print_green(initial_day)
        day, date, hour = parse_button(b["title"])
        
        aux_day = day
        colors.print_green(aux_day)
        print(date)
        print(hour)
        print("---------")


if __name__ == '__main__':
    BLINDEX = "https://darturnos.com/turnos/turnero/4188"
    CERRADA = "https://darturnos.com/turnos/turnero/4189"

    buttons = get_turnos(BLINDEX)
   