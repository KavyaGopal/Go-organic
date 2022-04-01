import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validator, Validators} from '@angular/forms';
import { GetApiService } from './../../get-api.service';
import { HttpClient} from '@angular/common/http';
import { Router} from '@angular/router';
import { CustomValidationService } from '../../services/custom-validation.service';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  constructor(private http: HttpClient, private router: Router, private customValidator: CustomValidationService, private fb: FormBuilder) {}

  
  loginForm = this.fb.group({
    email: new FormControl('', Validators.required),
    name: new FormControl('', Validators.required),
    phone: new FormControl('', Validators.required),
    address: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
    password2: new FormControl('', Validators.required)
  }, {
    validator: this.customValidator.passwordMatchValidator("password", "password2")
  })

  ngOnInit(): void {
  }

  get email() {
    return this.loginForm.get("email")
  }

  get password() {
    return this.loginForm.get("password")
  }

  get password2() {
    return this.loginForm.get("password2")
  }

  get name() {
    return this.loginForm.get("name")
  }
  get address() {
    return this.loginForm.get("address")
  }
  get phone() {
    return this.loginForm.get("phone")
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
      complete: () => {
      // alert("Login Success");
      this.loginForm.reset();
      this.router.navigate(['login'])
      }
    })
  }

}
