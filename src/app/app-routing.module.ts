import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomeComponent } from './components/home/home.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { CartComponent } from './components/cart/cart.component';
import { CheckoutComponent } from './components/checkout/checkout.component';
import { LoginComponent } from './components/login/login.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { TestimonialComponent } from './components/testimonial/testimonial.component';
import { PaymentComponent } from './components/payment/payment.component';


const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'home' },
  { path: 'home', component: HomeComponent  },
  { path: 'dashboard', component: DashboardComponent  },
  { path: 'cart', component: CartComponent  },
  { path: 'checkout', component: CheckoutComponent},
  { path: 'login', component: LoginComponent  },
  { path: 'sign-up', component: SignUpComponent  },
  { path: 'testimonial', component: TestimonialComponent  },
  { path: 'payment', component: PaymentComponent  }


  // { path: 'error', component: ErrorComponent },
  // { path: 'error/404', component: ErrorComponent },
  // { path: '**', redirectTo: 'error/404' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
