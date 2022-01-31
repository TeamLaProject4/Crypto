import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewwalletComponent } from './newwallet.component';

describe('NewwalletComponent', () => {
  let component: NewwalletComponent;
  let fixture: ComponentFixture<NewwalletComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewwalletComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewwalletComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
