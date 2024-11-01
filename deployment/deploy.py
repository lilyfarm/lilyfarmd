#!/usr/bin/env python3

import os

def main():
    msg = """Enter your USDA Farmers' Market API key below. If you do not have one, 
             you can get one at https://www.usdalocalfoodportal.com/fe/datasharing/"""
    print(msg)

    usda_credentials = input("Enter USDA Farmers' Market API Key:")

    msg = """Enter your geonames.org username. If you do not have one,
            you can get one at http://www.geonames.org/login."""
    print(msg)

    geonames_credentials = input("Enter geonames.org username:")

    msg = """Enter path to TLS certificate. If you do not have one, 
            you can get one using certbot."""
    print(msg)

    certificate = input("Enter TLS certificate path:")

    msg = """Enter path to TLS private key. If you do not have one,
             you can get one using certbot."""
    print(msg)

    key = input("Enter TLS private key path:")

    lilyfarmd =   f"sudo LILYFARM_USDA_CREDENTIALS={usda_credentials} " + \
                  f"LILYFARM_GEONAMES_CREDENTIALS={geonames_credentials} " + \
                  f"LILYFARM_TLS_CERTIFICATE={certificate} " + \
                  f"LILYFARM_TLS_KEY={key} " + \
                  "/usr/bin/lilyfarm"
    
    print("Writing service script with credentials to /usr/bin/lilyfarmd.sh")
    os.system("echo !/bin/bash | sudo tee /usr/bin/lilyfarmd.sh")
    os.system(f"echo {lilyfarmd} | sudo tee -a /usr/bin/lilyfarmd.sh")
    print("Done.")

    print("Giving the script execution permissions.")
    os.system(f"sudo chmod +x /usr/bin/lilyfarmd.sh")
    print("Done.")

    print("\nCopying lilyfarmd.service to /etc/systemd/system/")
    os.system("sudo cp lilyfarmd.service /etc/systemd/system/")
    print("Done.\n")

    print("\nCopying the lilyfarm executable from binaries to /usr/bin/")
    os.system("sudo cp ../binaries/lilyfarm /usr/bin/lilyfarm")
    os.system("sudo chmod +x /usr/bin/lilyfarm")
    print("Done.\n")

    print("Reloading the systemd manager.")
    os.system("sudo systemctl daemon-reload")

    print("Starting the service")
    os.system("sudo systemctl start lilyfarmd.service")

    print("\n\nAll done!")

if __name__ == '__main__':
    main()


