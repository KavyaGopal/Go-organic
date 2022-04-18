import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PaypalpaymentComponent } from './paypalpayment.component';

describe('PaypalpaymentComponent', () => {
  let component: PaypalpaymentComponent;
  let fixture: ComponentFixture<PaypalpaymentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PaypalpaymentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PaypalpaymentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
