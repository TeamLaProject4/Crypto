import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CryptieComponent } from './cryptie/cryptie.component';
import { LandingComponent } from './landing/landing.component';
import { LedgerComponent } from './ledger/ledger.component';
import { LogoutComponent } from './logout/logout.component';
import { NewwalletComponent } from './newwallet/newwallet.component';
import { PaymentComponent } from './payment/payment.component';
import { StakeComponent } from './stake/stake.component';
import { WalletComponent } from './wallet/wallet.component';

const routes: Routes = [
  { path: 'cryptie', component: CryptieComponent },
  { path: 'wallet', component: WalletComponent },
  { path: 'payment', component: PaymentComponent },
  { path: 'payment/:amount/:to', component: PaymentComponent },
  { path: 'ledger', component: LedgerComponent },
  { path: 'newwallet', component: NewwalletComponent },
  { path: 'stake', component: StakeComponent },
  { path: 'logout', component: LogoutComponent },
  { path: '**', component: LandingComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
