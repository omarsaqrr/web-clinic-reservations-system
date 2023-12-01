import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  
})
export class SignupComponent implements OnInit{
  userlist:any;

  ngOnInit(): void {
      
  }
  constructor(private http: HttpClient){
    this.userlist=[];
  }

 /* signup()
  {
    return this.http.get("").subscribe(result:any=>{
      this.userlist=result;
    });
  }
  */
  role: string = '';

  setUserType(selectedType: string) {
    this.role = selectedType;
  }

}
