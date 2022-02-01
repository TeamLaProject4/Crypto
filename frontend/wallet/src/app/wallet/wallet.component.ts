import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AbstractControl, FormBuilder, FormControl, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss']
})


export class WalletComponent  {
  public userForm:FormGroup; // variable is created of type FormGroup is created
  accountnumber: String = "";
  to_amount: String = "";

  constructor(  private fb: FormBuilder,
                private router: Router) {
    // Form element defined below
    this.userForm = this.fb.group({
      accountnumber: '',
      to_amount: ''
    });
  }

  makePayment() {
    this.to_amount=this.userForm.get('to_amount')?.value;
    this.accountnumber=this.userForm.get('accountnumber')?.value;
    this.router.navigate(["payment"], {
      queryParams: { amount: this.to_amount, to: this.accountnumber }
    }
    );
  }

  }
