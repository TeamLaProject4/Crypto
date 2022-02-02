import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-cryptie',
  templateUrl: './cryptie.component.html',
  styleUrls: ['./cryptie.component.scss'],
})
export class CryptieComponent implements OnInit {
  public userForm: FormGroup; // variable is created of type FormGroup is created

  public CryptieQRcode: string = '';
  public to_amount: string = '';
  public accountnumber: string = '';

  constructor(private fb: FormBuilder, private router: Router) {
    this.userForm = this.fb.group({
      accountnumber: '',
      to_amount: '',
    });
    this.CryptieQRcode = '_';
  }

  ngOnInit(): void {}

  generateQRcode() {
    this.to_amount = this.userForm.get('to_amount')?.value;
    this.accountnumber = this.userForm.get('accountnumber')?.value;
    this.CryptieQRcode =
      'localhost:4200/payment?to=' +
      this.accountnumber +
      '&amount=' +
      this.to_amount;
    this.router.navigate(['cryptie'], {
      queryParams: { amount: this.to_amount, to: this.accountnumber },
    });
  }
}
