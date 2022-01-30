import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CryptieComponent } from './cryptie/cryptie.component';
import { WalletComponent } from './wallet/wallet.component';
import { PaymentComponent } from './payment/payment.component';
import { LedgerComponent } from './ledger/ledger.component';
import { LandingComponent } from './landing/landing.component';
import { HttpClientModule } from '@angular/common/http';


@NgModule({
  declarations: [
    AppComponent,
    CryptieComponent,
    WalletComponent,
    PaymentComponent,
    LedgerComponent,
    LandingComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
