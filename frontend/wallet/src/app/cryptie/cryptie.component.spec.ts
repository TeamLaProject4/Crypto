import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CryptieComponent } from './cryptie.component';

describe('CryptieComponent', () => {
  let component: CryptieComponent;
  let fixture: ComponentFixture<CryptieComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CryptieComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CryptieComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
