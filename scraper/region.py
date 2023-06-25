from selenium.webdriver import Chrome
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import json
import os


def scrape_regions(region_url, province_id):
  regions = []
  region_driver = Chrome()
  region_driver.get(region_url)

  # wait for the javascript to be executed
  wait = WebDriverWait(region_driver, 10)
  wait.until(EC.visibility_of_element_located((By.CLASS_NAME, "header_mentok")))

  table_header = region_driver.find_element(By.CLASS_NAME, "header_mentok")
  data_table = table_header.find_element(By.XPATH, "following-sibling::tbody")
  rows = data_table.find_elements(By.TAG_NAME, "tr")

  pagination_links = []
  pagination_elm = table_header.find_element(By.XPATH, "..").find_element(By.XPATH, "following-sibling::center")
  pagination_link_elms = pagination_elm.find_elements(By.CLASS_NAME, "tpage")
  for pagination_link_elm in pagination_link_elms:
    pagination_link = pagination_link_elm.get_attribute("href")
    pagination_links.append(pagination_link)

  for row in rows:
    row_datas = row.find_elements(By.TAG_NAME, "td")

    village_name = row_datas[2].text
    subdistrict_name = row_datas[4].text
    city_status = row_datas[5].text
    city_name = row_datas[6].text
    
    regions.append({
      "city_name": f"{city_status} {city_name}",
      "subdistrict_name": subdistrict_name,
      "village_name": village_name
    })

  region_driver.close()

  for pagination_link in pagination_links:
    scrape_region(pagination_link, regions)

  write_regions_to_file(regions, province_id)

def scrape_region(region_url, regions):
  region_driver = Chrome()
  region_driver.get(region_url)

  # wait for the javascript to be executed
  wait = WebDriverWait(region_driver, 10)
  wait.until(EC.visibility_of_element_located((By.CLASS_NAME, "header_mentok")))

  table_header = region_driver.find_element(By.CLASS_NAME, "header_mentok")
  data_table = table_header.find_element(By.XPATH, "following-sibling::tbody")
  rows = data_table.find_elements(By.TAG_NAME, "tr")

  for row in rows:
    row_datas = row.find_elements(By.TAG_NAME, "td")

    village_name = row_datas[2].text
    subdistrict_name = row_datas[4].text
    city_status = row_datas[5].text
    city_name = row_datas[6].text
    
    regions.append({
      "city_name": f"{city_status} {city_name}",
      "subdistrict_name": subdistrict_name,
      "village_name": village_name
    })

  region_driver.close()

def write_regions_to_file(datas, province_id):
  regions = {}

  # mapping datas to regions, clean duplicate city and subdistrict
  for data in datas:
    city_name = data["city_name"]
    subdistrict_name = data["subdistrict_name"]
    village_name = data["village_name"]

    if city_name not in regions:
      regions[city_name] = {}

    if subdistrict_name not in regions[city_name]:
      regions[city_name][subdistrict_name] = []

    regions[city_name][subdistrict_name].append(village_name)

  # sort regions data
  regions = dict(sorted(regions.items(), key=lambda x: x[0]))
  for city_name, subdistricts in regions.items():
    subdistricts = dict(sorted(subdistricts.items(), key=lambda x: x[0]))
    regions[city_name] = subdistricts
    for subdistrict_name, villages in subdistricts.items():
      villages = sorted(villages)
      regions[city_name][subdistrict_name] = villages

  # assign id to regions
  city_id = 0
  regions_with_id = []
  for city_name, subdistricts in regions.items():
    city_id += 1
    region = {
      "id": city_id,
      "name": city_name,
      "subdistricts": []
    }
    subdistrict_id = 0
    for subdistrict_name, villages in subdistricts.items():
      subdistrict_id += 1
      subdistrict = {
        "id": subdistrict_id,
        "name": subdistrict_name,
        "villages": []
      }
      village_id = 0
      for village_name in villages:
        village_id += 1
        subdistrict["villages"].append({
          "id": village_id,
          "name": village_name,
        })
      region["subdistricts"].append(subdistrict)
    regions_with_id.append(region)

  # write to file
  directory_path = "data/cities"
  isExist = os.path.exists(directory_path)
  if not isExist:
    os.makedirs(directory_path)

  directory_path = "data/subdistricts"
  isExist = os.path.exists(directory_path)
  if not isExist:
    os.makedirs(directory_path)

  directory_path = "data/villages"
  isExist = os.path.exists(directory_path)
  if not isExist:
    os.makedirs(directory_path)

  cities = []
  print(f">>>> province: [{province_id}]")
  for city in regions_with_id:
    print({
      "id": city["id"],
      "name": city["name"]
    })
    cities.append({
      "id": city["id"],
      "name": city["name"]
    })
    subdistricts = []

    for subdistrict in city["subdistricts"]:
      subdistricts.append({
        "id": subdistrict["id"],
        "name": subdistrict["name"]
      })
      filename = f"data/villages/{province_id}_{city['id']}_{subdistrict['id']}.json"
      with open(filename, "w") as file:
          json.dump(subdistrict["villages"], file)

    filename = f"data/subdistricts/{province_id}_{city['id']}.json"
    with open(filename, "w") as file:
        json.dump(subdistricts, file)

  filename = f"data/cities/{province_id}.json"
  with open(filename, "w") as file:
      json.dump(cities, file)
