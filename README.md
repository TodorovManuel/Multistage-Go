-- COMO GENERAR LA IMAGEN --
Necesitamos el archivo Dockerfile con su respectivo codigo
Luego "buildeamos la imagen con este comando" => 'docker build -t nombreDeTuImagen .'          (en nuestro gaso goapp) 

-- COMO CORRER LA IMAGEN --
Una vez que se "buildeo" sin problemas hay que ejecutar el siguiente comando 
'docker run -p 3000:3000 nombreDeTuImagen'
