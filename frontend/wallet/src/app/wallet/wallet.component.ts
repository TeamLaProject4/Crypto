import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {
  AbstractControl,
  FormBuilder,
  FormControl,
  FormGroup,
} from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss'],
})
export class WalletComponent {
  public userForm: FormGroup; // variable is created of type FormGroup is created
  accountnumber: string = '';
  to_amount: string = '';
  balance: string = '';
  PK: string = '';

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private api: APIService
  ) {
    // Form element defined below
    this.userForm = this.fb.group({
      accountnumber: '',
      to_amount: '',
    });
  }

  ngOnInit() {
    this.getAccountnumber();
    this.loadBalance();
  }

  loadBalance() {
    this.api.getBalance(this.accountnumber).subscribe((data) => {
      this.balance = data;
    });
  }

  getAccountnumber() {
    this.api.getPublicKey().subscribe((data) => {
      this.accountnumber = data;
    });
  }

  makePayment() {
    this.to_amount = this.userForm.get('to_amount')?.value;
    this.accountnumber = this.userForm.get('accountnumber')?.value;
    this.router.navigate(['payment'], {
      queryParams: { amount: this.to_amount, to: this.accountnumber },
    });
  }
}
