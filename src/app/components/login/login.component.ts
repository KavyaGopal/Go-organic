import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient} from '@angular/common/http';
import { Router} from '@angular/router';
import { CartService } from '../../services/cart.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  login = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required])
  });

  constructor(private http: HttpClient, private router: Router, private cartService: CartService) {}

  ngOnInit(): void {
  }

  get email() {
    return this.login.get('email')
  }

  get password() {
    return this.login.get("password")
  }

  onSubmit() {
    // console.warn(this.login.value)
    this.http.post<any>("http://localhost:8000/loginUser", this.login.value)
    .subscribe({
      next: (value) => localStorage.setItem('userData', value.name),
      error: (e) => {alert("Something went wrong");
      console.log("The error is:");
      console.warn(e)
    },
      complete: () => {
      // alert("Login Success");
      this.login.reset();
      this.cartService.userName.next(localStorage.getItem('userData'));
      // console.warn(localStorage.getItem('userData'));
      this.router.navigate(['home']);

      }
    })
  }
}
