# Hecate - Tools

Arsenal is the marketplace for tools. It is a collection of tools that can be used for different purposes. The tools are classified by categories and subcategories. The tools can be used for different purposes such as penetration testing, security assessments, web application security testing, network security, etc.

## Tools

### Organisation of tools.json

Exemple :
```json
        {
            "id": "1",
            "name": "nmap",
            "description": "Network exploration tool and security/port scanner.",
            "version": "7.91",
            "author": "Fyodor",
            "commands": ["nmap -v -A scanme.nmap.org"],
            "actions": {
                "install": "apt-get install nmap",
                "remove": "apt-get remove nmap",
                "update": "apt-get update nmap"
            },
            "metadata": {
                "license": "GPL",
                "source": "https://nmap.org/"
            }
        }
```
