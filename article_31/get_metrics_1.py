import requests

response = requests.get("http://localhost:8080/debug/vars")

assert response.status_code == 200, "Chyba při čtení metrik: neočekávaný HTTP kód odpovědi"

payload = response.json()

print("Počet překreslení: ", payload["renderingCounter"])
print("Celkový čas výpočtu: ", payload["humanReadableDuration"])
print("Počet obarvených pixelů: ", payload["pixelsColored"])
