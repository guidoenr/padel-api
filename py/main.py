import requests
import colors
import sys
from bs4 import BeautifulSoup
from datetime import datetime

BLINDEX = "https://darturnos.com/turnos/turnero/4188"
CERRADA = "https://darturnos.com/turnos/turnero/4189"


class Day():
    def __init__(self, day_name:str, date:str):
        self.day_name = day_name
        self.date = date
        self.turnos = []

    def show_turnos(self):
        colors.print_green(f'{self.day_name} - {self.date}')
        if len(self.turnos) == 0:
            colors.print_red("SIN TURNOS")
        for hour in self.turnos:
            colors.print_bold(hour)
        print("---------------------")

    def add_turno(self, hour:str):
        self.turnos.append(hour)


def parse_button(button):
    splitted = button["title"].split()
    date = splitted[5]
    hour = splitted[8]

    # transform the date string into a datetime format
    date_time_obj = datetime.strptime(date, '%d/%m/%Y')
    
    day_name = date_time_obj.strftime('%A').upper()
    
    return day_name, date, hour

def get_turnos(field):
    page = requests.get(field)
    soup = BeautifulSoup(page.content, "html.parser")
    return soup.find_all("button") # buttons are the Tags that contains all the info about a turno
    

def save_turnos(field):
    days = []
    
    # all the turnos 
    turnos_cancha = get_turnos(field)

    # first turno
    aux_day, date, hour = parse_button(turnos_cancha[0])

    # first day
    new_day = Day(aux_day, date)

    for turno in turnos_cancha:
        day, date, hour = parse_button(turno)
        if aux_day == day:
            new_day.add_turno(hour)
        else:
            days.append(new_day)
            new_day = Day(day, date)
            new_day.add_turno(hour)
            aux_day = day
    
    for d in days:
        d.show_turnos()

    
        
if __name__ == '__main__':
    field = ""
    cancha_input = sys.argv[1]
    if cancha_input == "blindex":
        field = BLINDEX
        colors.print_title("CANCHA BLINDEX")
    else:
        field = CERRADA
        colors.print_title("CANCHA CERRADA")

    save_turnos(field)

   