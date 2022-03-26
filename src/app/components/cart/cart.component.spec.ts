import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { CartComponent } from './cart.component';

describe('CartComponent', () => {
  let component: CartComponent;
  let fixture: ComponentFixture<CartComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CartComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('Checkout button should be enabled', () => {
    // expect(h1.textContent).toContain(component.title);
    let button = fixture.debugElement.nativeElement.querySelector('#checkoutButton');
    expect(button.disabled).toBeFalsy();
  });

  it('Check Button name', () => {
    // expect(h1.textContent).toContain(component.title);
    let button = fixture.debugElement.nativeElement.querySelector('#checkoutButton');
    expect(button.innerHTML).toBe('Checkout');
  });

  // it('Check Table Columns', () => {
  //   // expect(h1.textContent).toContain(component.title);
  //   let title = fixture.debugElement.nativeElement.querySelector('#totalID');
  //   expect(title.innerHTML).toBe('Total');
  //   console.log(title.innerHTML)
  // });

});
