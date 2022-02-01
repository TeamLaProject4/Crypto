import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AbstractControl, FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';
import {RouterModule} from '@angular/router';
@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.scss']
})
export class LandingComponent implements OnInit {

  public loginForm:FormGroup;
  seedphrase: string = "";

  constructor(  private fb: FormBuilder,
                private router: Router,
                private http: HttpClient,
                private api: APIService) {
    // Form element defined below
    this.loginForm = this.fb.group({
      seedphrase: '',
    });
  }

  ngOnInit(): void {
  }


  login(){
    this.seedphrase=this.loginForm.get('seedphrase')?.value;
    this.confirmWallet(this.seedphrase);
    this.router.navigate(["wallet"], {
    });
  }

  confirmWallet(seedphrase: string){
    this.api.confirmMnemonic(seedphrase)
    .subscribe(data => {
      console.log(data);
    });
  }
  }
