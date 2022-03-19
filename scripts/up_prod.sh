#!/bin/bash
sudo docker-compose -f prod.yml up --build -d --scale backend=5