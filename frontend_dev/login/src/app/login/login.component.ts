import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  
})

export class LoginComponent implements OnInit{
  email: string ;
  password: string ;
  role: string;
  isSuccessful: Boolean;

  ngOnInit(): void {
      
  }
  constructor(private http: HttpClient,private router: Router){

    this.email= '';
    this.password = '';
    this.role="";
    this.isSuccessful=false;
  
  }

 login() {
  const apiUrl = 'http://localhost:8080';
  const signinEndpoint = '/sign-in';


  const requestBody = {
    email: this.email,     
    password: this.password, 
  };

  
  this.http.post(apiUrl + signinEndpoint, requestBody).subscribe(
    (result: any) => {
      this.isSuccessful =true;
      this.setRole(result['object']['Role']);
      // Handle successful response
      console.log('SignIn successful:', result);
      
     
    },
    (error) => {
      // Handle error response
  
      console.error('Signup error:', error);
    }
  );
}

 

 
  setEmail(selectedType: string) {
    this.email = selectedType;
  }
  setPassword(selectedType: string) {
    this.password= selectedType;
  }
  setRole(selectedType: string) {
    this.role = selectedType;
  }
  
}