import { Injectable } from '@angular/core';
import {Subject} from 'rxjs'

@Injectable({
  providedIn: 'root'
})
export class CartService {

  constructor() { }
  cartSubject = new Subject<any>();
  userName = new Subject<any>();
  
  getUser(){
    var user = null;
    
    if (localStorage.getItem('userData') != null) {
      user = localStorage.getItem('userData');
      
  }
  return user;
  }
}