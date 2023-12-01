# gin-starter
Gin starter project, with mysql, gorm, viper, nginx, and docker

You can clone the project and run command:
```bash
sudo docker compose up -d
```
So docker pulls the images, and serves mysql, project and nginx. 
You can access the project through your explorer and address [http://localhost](http://localhost).

Also you can change exposed ports, volumes of the containers, and the rest configs as you need.
Environment file `.env` located in `/cmd/server/` directory, so you can change env keys there.

Then start developing your idea...

Any comment or contribution would be appreciated :)
