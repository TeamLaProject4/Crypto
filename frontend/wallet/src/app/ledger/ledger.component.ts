import { HttpClient } from '@angular/common/http';
import { temporaryAllocator } from '@angular/compiler/src/render3/view/util';
import {
  AfterViewInit,
  Component,
  Pipe,
  ViewChild,
  PipeTransform,
  NgModule,
} from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource, MatTable } from '@angular/material/table';
import { ActivatedRoute, Router } from '@angular/router';
import { APIService } from '../api.service';
import { FormBuilder, FormGroup } from '@angular/forms';
import { CustomSlice } from '../pipes/customslice.pipe';

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
  accountnumber: string = '';
  full_ledger: string = '';
  public accountnumberform: FormGroup; // variable is created of type FormGroup is created

  constructor(
    private http: HttpClient,
    private api: APIService,
    private route: ActivatedRoute,
    private router: Router,
    private fb: FormBuilder
  ) {
    this.accountnumberform = this.fb.group({
      accountnumber: '',
    });

    this.getLedgerEntries(this.accountnumber);
    this.dataSource = new MatTableDataSource(this.entries);

    console.log(this.entries);
  }

  ngOnInit() {
    // this.route.queryParams.subscribe((params) => {
    //   console.log(params['accountnumber']);
    //   this.accountnumber = params['accountnumber'];
    // });
    this.getQueryParams();
    this.accountnumber = this.accountnumberform.get('publicKey')?.value;
  }

  searchAccountnumber(accountnumber: string) {
    this.accountnumber = accountnumber;
    this.getLedgerEntries(this.accountnumber);
  }

  getQueryParams() {
    this.route.queryParams.subscribe((params) => {
      this.accountnumber = params['accountnumber'];
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
    console.log(accountnumber);
    this.api.getTransactions(accountnumber).subscribe((data) => {
      data.forEach((entry: any) => {
        const temp: LedgerEntry = {
          sender: entry.sender_pk.substring(0, 5),
          amount: entry.amount,
          receiver: entry.receiver_pk.substring(0, 5),
        };
        this.entries.push(temp);
      });

      this.dataSource = new MatTableDataSource(this.entries);
      this.refresh();
    });
  }

  refresh() {
    this.dataSource = new MatTableDataSource(this.entries); //this.response
    this.dataSource.sort = this.sort;
    console.log(this.dataSource);
    // console.log("hoi");
  }
  // }
}
