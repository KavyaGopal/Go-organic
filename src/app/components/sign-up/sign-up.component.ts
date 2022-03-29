import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validator, Validators} from '@angular/forms';
import { GetApiService } from './../../get-api.service';
import { HttpClient} from '@angular/common/http';
import { Router} from '@angular/router';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  constructor(private http: HttpClient, private router: Router) {}

  
  loginForm = new FormGroup({
    email: new FormControl('', Validators.required),
    userName: new FormControl('', Validators.required),
    phNumber: new FormControl('', Validators.required),
    address: new FormControl('', Validators.required),
    password1: new FormControl('', Validators.required),
    password2: new FormControl('', Validators.required)
  })

  ngOnInit(): void {
  }

  get email() {
    return this.loginForm.get("email")
  }

  get password1() {
    return this.loginForm.get("password1")
  }
  get password2() {
    return this.loginForm.get("password2")
  }
  get userName() {
    return this.loginForm.get("userName")
  }
  get address() {
    return this.loginForm.get("address")
  }
  get phNumber() {
    return this.loginForm.get("phNumber")
  }

  onSubmit() {
    console.warn(this.loginForm.value)
    this.http.post<any>("http://localhost:8000/registerUser", this.loginForm.value)
    .subscribe({
      next: (value) => console.log(value),
      error: (e) => {alert("Something went wrong");
      console.log("The error is:");
      console.warn(e)
    },
      complete: () => {alert("Login Success");
      this.loginForm.reset();
      // this.router.navigate(['login'])
      }
    })
  }

}
