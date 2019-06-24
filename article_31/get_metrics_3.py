import requests
import os
import csv

filename = "mandelbrot.csv"

response = requests.get("http://localhost:8080/debug/vars")

assert response.status_code == 200, "Chyba při čtení metrik: neočekávaný HTTP kód odpovědi"

payload = response.json()

mode = None
write_header = False

if os.path.exists(filename):
    mode = "a"
else:
    mode = "w"
    write_header = True

with open(filename, mode) as csv_file:
    writer = csv.writer(csv_file, delimiter=',')
    if write_header:
        writer.writerow(("Počet překreslení", "Čas výpočtu", "Obarvených pixelů", "Poslední výpočet"))
    writer.writerow((payload["renderingCounter"],
                     payload["renderingDuration"],
                     payload["pixelsColored"],
                     payload["humanReadableDuration"]))
