# Instagram Backend API made in GO - Jasjeet Singh Siddhu

## Mongo DB Database Structure-

<img src="https://user-images.githubusercontent.com/68146112/136667867-e81ac4f4-a433-447c-9d0c-f37d326d1a69.png">

## Initial Collections are empty -

<img src="https://user-images.githubusercontent.com/68146112/136667932-8ccd64d4-0cef-4963-9519-86c484b792cd.png">

<img src="https://user-images.githubusercontent.com/68146112/136667956-3bbfe89c-128e-4850-940d-28e3a644f931.png">
<hr>

## Task 1 - Psswords r hashed
/users
### Frontend Form -
<img src="https://user-images.githubusercontent.com/68146112/136668068-09d3905b-0e62-47ad-a310-b76fe9c388fa.png">

### Submitting Form-

<img src="https://user-images.githubusercontent.com/68146112/136668132-4dc76966-76d0-4cad-afcc-0205c4267dfa.png">

### On Submission -
Frontend displays-
<img src="https://user-images.githubusercontent.com/68146112/136668189-a850ff21-b678-4447-922f-b16924f5bee2.png">

MongoDB entry -

Password is hashed
<img src="https://user-images.githubusercontent.com/68146112/136668261-df62fba7-1eb4-4f27-b18d-150d32367b3e.png">
<hr>

## Task 2
put object id of mongodb as id

/users/id

### Users in Database -
<img src="https://user-images.githubusercontent.com/68146112/136668378-e18aed0c-c33b-4ed1-9d55-0382e053aa40.png">

### Now get request by user id -
/users/?id=6161cfe51fa3e67f4df735b2
<img src="https://user-images.githubusercontent.com/68146112/136668451-5f5b10d0-8c57-4f87-b223-0eda299290fd.png">
User details r displayed 

### if we pass POST request -
<img src="https://user-images.githubusercontent.com/68146112/136668493-043c37fe-880c-4eda-acfa-35c238b59416.png">
As only GET request is allowed
<hr>

## Task 3 
/posts
### Frontend Form -
<img src="https://user-images.githubusercontent.com/68146112/136668603-d5142b93-426a-490b-b5d4-fdac3d5a4e19.png">

### Submitting Form-

<img src="https://user-images.githubusercontent.com/68146112/136668635-3f262940-1789-4bbd-9eff-7f941f309572.png">

### On Submission -
Frontend displays-
<img src="https://user-images.githubusercontent.com/68146112/136668646-75903ecd-7da3-4792-af24-5d7764cdad73.png">

### Images r added in project directory after form submission automatically
<img src="https://user-images.githubusercontent.com/68146112/136669255-7108ceb7-b782-4dd3-99b4-bdfefa98d5e8.png">

Local Terminal -
<img src="https://user-images.githubusercontent.com/68146112/136668690-15b8427f-06e2-4da9-915e-4870191162b9.png">

MongoDB entry -


<img src="https://user-images.githubusercontent.com/68146112/136668721-d091bd69-16d7-45f5-9c53-c11933178498.png">
<hr>

## Task 4
put object id of post in mongodb as id

/posts/id

### Posts in Database -
<img src="https://user-images.githubusercontent.com/68146112/136668911-aed8bc0b-1b70-4093-a83c-ecea84981323.png">

### Now get request by user id -
/posts/?id=6161d47c1fa3e67f4df735bc
<img src="https://user-images.githubusercontent.com/68146112/136668955-a69e8277-f6f1-4d35-abee-5c0ae1e83333.png">
Post details r displayed 

### if we pass POST request -
<img src="https://user-images.githubusercontent.com/68146112/136669001-f6787d8e-f1a6-40bf-9d1a-176d6fbba218.png">
As only GET request is allowed
<hr>

## Task 5

/posts/user/id

Enter user object id in Mongo db as id here

### /posts/users/?id=6161cfe51fa3e67f4df735b2

Should give 2 posts details as 2 posts are made by Kiwi -
<img src="https://user-images.githubusercontent.com/68146112/136669101-bead9987-f61b-4d4f-b48f-91b232ef2021.png">

### /posts/users/?id=6161cec91fa3e67f4df735b0

Should give 2 posts details as 2 posts are made by Jasjeet -
<img src="https://user-images.githubusercontent.com/68146112/136669144-a3da1f45-f42f-47df-aabb-5cf86dbef04a.png">

### If Post request sent-
<img src="https://user-images.githubusercontent.com/68146112/136669180-a2502ac6-6cc6-475c-a21d-ce3cf07ae7ab.png">
 As only get allowed
 <hr>

 ## Tasks completed -

 ### All api endpoints as per requirement made in GO using Standard libraries

### Backend database soley made in mongodb
 ### All Passwords hashed so that no reverse engineering can be done


