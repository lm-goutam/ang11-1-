import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
 

  Org: any[] = [
    { id: 1, name: 'Lemma' },
    { id: 2, name: 'Wipro' },
    { id: 3, name: 'Tcs' },
    { id: 4, name: 'Zoho' }
  ];
  Cms: any[] = [
    { id: 1, name: 'Board Sign' },
    { id: 2, name: 'Ayuda' },
    { id: 3, name: 'Xtream Media' },
    { id: 4, name: 'Novastar' },
    { id: 5, name: 'Vinnox' },
    { id: 6, name: 'OnSignTv' }
  ];
  Stat: any[] = [
    { id: 1, name: 'Done' },
    { id: 2, name: 'Pending' },
    { id: 3, name: 'Inprogress' },
    { id: 4, name: 'Blocked' }
  ];
}
