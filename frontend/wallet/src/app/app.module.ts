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
import { NewwalletComponent } from './newwallet/newwallet.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { QRCodeModule } from 'angularx-qrcode';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { LogoutComponent } from './logout/logout.component';
import { StakeComponent } from './stake/stake.component';
import { CustomSlice } from './pipes/customslice.pipe';

@NgModule({
  declarations: [
    AppComponent,
    CryptieComponent,
    WalletComponent,
    PaymentComponent,
    LandingComponent,
    NewwalletComponent,
    LedgerComponent,
    LogoutComponent,
    StakeComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    QRCodeModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatTableModule,
    MatPaginatorModule,
  ],
  providers: [CustomSlice],
  bootstrap: [AppComponent],
})
export class AppModule {}
