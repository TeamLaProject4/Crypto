import { HttpClient } from '@angular/common/http';
import { temporaryAllocator } from '@angular/compiler/src/render3/view/util';
import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { ActivatedRoute } from '@angular/router';
import { APIService } from '../api.service';

export interface LedgerEntry {
  sender: string;
  amount: string;
  receiver: string;
}

export interface Transactions {
  sender_pk: string;
  receiver_pk: string;
  amount: number;
  tx_type: number;
  id: string;
  timestamp: number;
  signature: string;
}

@Component({
  selector: 'app-ledger',
  templateUrl: './ledger.component.html',
  styleUrls: ['./ledger.component.scss'],
})
export class LedgerComponent implements AfterViewInit {
  displayedColumns: string[] = ['sender', 'amount', 'receiver'];
  dataSource: MatTableDataSource<LedgerEntry>;

  @ViewChild(MatPaginator)
  paginator!: MatPaginator;
  @ViewChild(MatSort)
  sort!: MatSort;

  entries: LedgerEntry[] = [];
  accountNumber: string = '';
  full_ledger: string = '';

  constructor(
    private http: HttpClient,
    private api: APIService,
    private route: ActivatedRoute
  ) {
    this.getAccountNumber();
    this.getLedgerEntries(this.accountNumber);
    this.dataSource = new MatTableDataSource(this.entries);

    console.log(this.entries);
  }

  getAccountNumber() {
    this.route.queryParams.subscribe((params) => {
      this.accountNumber = params['accountNumber'];
    });
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

  /** gets the ledger entries tied to the accountnumber */
  getLedgerEntries(accountnumber: string) {
    this.api.getTransactions(accountnumber).subscribe((data) => {
      data.forEach((entry: any) => {
        const temp: LedgerEntry = {
          sender: entry.sender_pk,
          amount: entry.amount,
          receiver: entry.receiver_pk,
        };
        this.entries.push(temp);
      });
    });
  }
}
