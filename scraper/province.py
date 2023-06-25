from selenium.webdriver import Chrome
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import json
import os

from region import scrape_regions

def scrape_provinces():
  province_driver = Chrome()
  province_driver.get('https://www.nomor.net')

  # wait for the javascript to be executed
  wait = WebDriverWait(province_driver, 10)
  wait.until(EC.visibility_of_element_located((By.CLASS_NAME, "header_mentok")))

  provinces = []
  table_header = province_driver.find_element(By.CLASS_NAME, "header_mentok")
  data_table = table_header.find_element(By.XPATH, "following-sibling::tbody")
  rows = data_table.find_elements(By.TAG_NAME, "tr")

  province_count = 1
  for row in rows:
    if province_count == len(rows):
      # skip the last row that contain the total datas (which we don't need)
      break

    row_datas = row.find_elements(By.TAG_NAME, "td")
    province_elm = row_datas[1].find_element(By.TAG_NAME, "a")
    province_name = province_elm.text
    region_url = row_datas[9].find_element(By.TAG_NAME, "a").get_attribute("href")
    provinces.append({
      "id": province_count,
      "province_name": province_name
    })
    scrape_regions(region_url, province_count)
    province_count += 1
    
  write_provinces_to_file(provinces)

  province_driver.close()

def write_provinces_to_file(provinces):
  directory_path = "data"
  isExist = os.path.exists(directory_path)
  if not isExist:
    os.makedirs(directory_path)

  filename = "data/provinces.json"
  with open(filename, "w") as file:
      json.dump(provinces, file)