import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
// import { FlexLayoutModule } from '@angular/flex-layout';
import { MaterialModule } from './material/material.module';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LayoutModule } from '@angular/cdk/layout';
import { SlickCarouselModule } from 'ngx-slick-carousel';

import { NavigationHeaderComponent } from './components/navigation-header/navigation-header.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { HomeComponent } from './components/home/home.component';
import { CartComponent } from './components/cart/cart.component';
import { TestimonialComponent} from "./components/testimonial/testimonial.component"
import{ HttpClientModule,  HttpClient} from '@angular/common/http';
import { CheckoutComponent } from './Components/checkout/checkout.component';
import { LoginComponent } from './components/login/login.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatFormFieldModule} from '@angular/material/form-field';
import { MatInputModule} from '@angular/material/input';
import { MatButtonModule} from '@angular/material/button';
import { MatCardModule} from '@angular/material/card';
import {MatToolbarModule} from '@angular/material/toolbar';
import { ReactiveFormsModule} from '@angular/forms';
import { FormsModule} from '@angular/forms';
import { PaypalpaymentComponent } from './components/paypalpayment/paypalpayment.component';
import { PaymentComponent } from './components/payment/payment.component';




@NgModule({
  declarations: [
    AppComponent,
    NavigationHeaderComponent,
    DashboardComponent,
    HomeComponent,
    TestimonialComponent,
    CartComponent,
    CheckoutComponent,
    LoginComponent,
    SignUpComponent,
    PaypalpaymentComponent,
    PaymentComponent
     
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    SlickCarouselModule,
    MaterialModule,
    HttpClientModule,
    LayoutModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatToolbarModule,
    ReactiveFormsModule,
    FormsModule
  ],
  providers: [HttpClient],
  bootstrap: [AppComponent]
})
export class AppModule { }
