// api.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  signup(email: string, password: string, role: string): Observable<any> {
    const signupEndpoint = role === 'doctor' ? '/sign-up/doctor' : '/sign-up/patient';

    const requestBody = {
      email: email,
      password: password,
      role: role,
    };

    return this.http.post<any>(this.apiUrl + signupEndpoint, requestBody).pipe(
      catchError((error) => {
        this.handleError(error);
        return throwError(error);
      })
    );
  }

  private handleError(error: any) {
    console.error('API Error:', error);
    // Perform common error-handling actions
    // ...
    // Example: Show a user-friendly error message
    // alert('An error occurred. Please try again later.');
  }
}
