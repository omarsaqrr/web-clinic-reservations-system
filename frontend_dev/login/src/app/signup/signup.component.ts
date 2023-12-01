import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  
})

export class SignupComponent implements OnInit{
  email: string ;
  password: string ;
  role: string ;
  isSuccessful: Boolean;

  ngOnInit(): void {
      
  }
  constructor(private http: HttpClient,private router: Router){

   this.email= '';
    this.password = '';
    this.role= '';
    this.isSuccessful=false;
  
  }

 signup() {
  const apiUrl = 'http://localhost:8080';
  const signupEndpoint = this.role === 'doctor' ? '/sign-up/doctor' : '/sign-up/patient';


  const requestBody = {
    email: this.email,     
    password: this.password, 
    role: this.role
  };

  this.http.post(apiUrl + signupEndpoint, requestBody).subscribe(
    (result: any) => {
      this.isSuccessful =true;
      // Handle successful response
      console.log('Signup successful:', result);
      
     
    },
    (error) => {
      // Handle error response
  
      console.error('Signup error:', error);
    }
  );
}

 

  setUserType(selectedType: string) {
    this.role = selectedType;
  }
  setEmail(selectedType: string) {
    this.email = selectedType;
  }
  setPassword(selectedType: string) {
    this.password= selectedType;
  }

}