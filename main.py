from ast import parse
from webbrowser import get
import requests
import colors
from bs4 import BeautifulSoup
from datetime import datetime

class Day():
    def __init__(self, day_name:str, date:str):
        self.day_name = day_name
        self.date = date
        self.turnos = []

    def show_turnos(self):
        colors.print_green(f'{self.day_name} - {self.date}')
        if len(self.turnos) == 0:
            colors.print_red("SIN TURNOS")
        for t in self.turnos:
            print(t)

    def add_turno(self, hour:str):
        self.turnos.append(hour)


def save_turnos():
    BLINDEX = "https://darturnos.com/turnos/turnero/4188"
    CERRADA = "https://darturnos.com/turnos/turnero/4189"
    days = []
    
    # all the turnos 
    turnos_cancha = get_turnos(BLINDEX)

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

        
def parse_button(button):
    splitted = button["title"].split()
    date = splitted[5]
    hour = splitted[8]

    date_time_obj = datetime.strptime(date, '%d/%m/%Y')
    day_name = date_time_obj.strftime('%A')
    
    return day_name, date, hour

def get_turnos(field):
    page = requests.get(field)
    soup = BeautifulSoup(page.content, "html.parser")
    return soup.find_all("button") # buttons are the Tags that contains all the info about a turno
    
        
if __name__ == '__main__':


    save_turnos()

   