# RecipeApplication
(AdigeMane) Golang Project

#### A fullstack application which serves as a social media for sharing recipes.
*Tech Stack: Golang (Go), HTML5/CSS, Bootstrap4, MongoDb*

![Home Page](/Screenshots/home_screen.jpg)


<details>
  <summary>Click here to see screen-shots of the application</summary>
  <img src="/Screenshots/options_screen.jpg" name="Menus">
  <img src="/Screenshots/recipe_login.jpg" name="Login Screen">
  <img src="/Screenshots/user_profile_screen.jpg" name="User profile screen">
  <img src="/Screenshots/recipe_info.jpg" name="Recipe Information">
  <img src="/Screenshots/add_recipe_screen.jpg" name="Add Recipe">
</details>

### API DOCUMENTATION:
* External APIs for the application:
  * **Sign-Up New User:**
  ``` 
  Request Fields:
      {
        username: string,
        password: password<string>
      }
  API Response:
      {
        name: username<string>,
        token: sessionToken<string>,
        result: success/failure<string>
      }
  ```

  * **Sign-in to the application:**
  ``` 
  Request Fields:
      {
        username: string,
        password: password<string>
      }
  API Response:
      {
        name: username<string>,
        token: sessionToken/nullValue<string>,
        result: success/failure<string>
      }
  ```
  * **Add User Info:**
  ``` 
  Request Fields:
      {
        username:username<string>,
        token:sessionToken<string>,
        location:userlocation<string>,
        about:userAbout<string>
      }
  API Response:
      {
        name: username<string>,
        token: sessionToken/nullValue<string>,
        result: success/failure<string>
      }
  ```
  * **Display Home Page:**
  ``` 
  Request Fields:
      {
        username:name<string>,
        token:sessionToken<string>
      }
  API Response:
      {
        username:name<string>,
        recipes:recipeArray[]<GetRecipe>,
        status:success/failure<string>
      }
  GetRecipe Struct:
      {
        username:name<string>,
        created:date<string>,
        recipename:recipeName<string>,
        recipedescription:recipeDescription<string>,
        recipesteps:recipeDirections<string>,
        uid:uniqueId<string>
      }
  ```
    * **Add new recipe:**
  ``` 
  Request Fields:
      {
        username:name<string>,
        token:sessionToken<string>,
        recipename:recipeName<string>,
        recipedescription:recipeDesc<string>,
        recipesteps:recipeInstruct<string>
      }
  API Response:
      {
        name: username<string>,
        token: sessionToken/nullValue<string>,
        result: success/failure<string>
      }
  Struct used to insert into database:
      {
        username:name<string>,
        created:date<string>,
        recipename:recipeName<string>,
        recipedescription:recipeDesc<string>,
        recipesteps:recipeInstruct<string>,
        uid:uniqueId<string>
      }
  ```
     * **Fetch and Display a particular recipe:**
  ``` 
  Request Fields:
      {
        uid:uniqueId<string>
      }
  API Response:
      returns a html page that displays the details of the recipe containing: Title, body, name, date, recipeName, recipeDescription, RecipeInstructions, and unique id.
  ```
     * **Delete a recipe:**
  ```
  Request Fields:
      {
        username:name<string>,
        token:sessionToken<string>,
        uid:uniqueId<string>
      }
   
  API Response:
      {
        name: username<string>,
        token: sessionToken/nullValue<string>,
        result: success/failure<string>
      }
  ```
       * **Load User Details:**
  ```
  Request Fields:
      {
        username:name<string>,
        token:sessionToken<string>
      }
   
  API Response:
      {
        name: username<string>,
        location: userLocation<string>,
        about: aboutUser<string>,
        result: success/failure<string>
      }
  ```
* Internal APIs for the application:
  * **Unique User Identifier handler to provide unique usernames across the app**
  ```
  Request Fields:
      {
        name:username<string>,
        valid:valid/invalid username<string>
      }
      
