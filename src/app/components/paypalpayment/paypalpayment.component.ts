import { render } from 'creditcardpayments/creditCardPayments'
import { Router } from '@angular/router'; // import router from angular router
import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';

@Component({
  selector: 'app-paypalpayment',
  templateUrl: './paypalpayment.component.html',
  styleUrls: ['./paypalpayment.component.css']
})

export class PaypalpaymentComponent implements OnInit {

  displayedColumns: string[] = ['id', 'itemName', 'imgSrc', 'itemCost', 'itemDesc',
    'itemSubtotal', 'star'];
  dataSource = new MatTableDataSource<cartItem>();
  getCartDetails: any = [];
  total: number = 0;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  constructor(private route: Router) {
    
    // TO DO: render takes JSON structure, we need to pass total amount to value as a string for render to send a proper request to PayPal API
    render(
      {
        id: "#myCheckoutPayPal",
        currency: "USD",
        value: (100).toString(),
        onApprove: (details) => {
          // alert("Transaction Successful. You may close the browser now.");
          this.route.navigate(['/payment']);

        }
      }
    );
  }

  ngOnInit(): void {
    this.cartDetails();
    this.getTotal();
  }

  cartDetails() {
    if (localStorage.getItem('cartData')) {
      this.getCartDetails = JSON.parse(localStorage.getItem('cartData'));
      console.log(this.getCartDetails);
      this.dataSource = new MatTableDataSource<cartItem>(this.getCartDetails);
    }
  }

  getTotal() {
    if (localStorage.getItem('cartData')) {
      this.getCartDetails = JSON.parse(localStorage.getItem('cartData'));
      this.total = this.getCartDetails.reduce((acc, val) => {
        return acc + (val.itemCost * val.itemQuantity) //add quantity in schema
      }, 0);
    }
  }

}

export interface cartItem {
  id: number;
  imgSrc: string;
  itemCost: number;
  itemDesc: string;
  itemGross: number;
  itemName: string;
  itemWt: number;

}