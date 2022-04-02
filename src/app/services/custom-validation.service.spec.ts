import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { CustomValidationService } from './custom-validation.service';

describe('CustomValidationService', () => {
  let service: CustomValidationService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CustomValidationService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
