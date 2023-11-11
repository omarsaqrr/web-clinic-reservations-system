import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { UserComponent } from './user/user.component';
import { DoctorComponent } from './doctor/doctor.component';

const routes: Routes = [

  {

    path: '',
  loadChildren:()=>import('./public/public.module').then((m)=>m.PublicModule),

}

,
{
  path:'sign up',
  component:SignupComponent
},
{
  path:'login',
  component:LoginComponent
}
,
{path:'user',
component:UserComponent
}
,
{
  path:'doctor'
  ,
  component: DoctorComponent
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
