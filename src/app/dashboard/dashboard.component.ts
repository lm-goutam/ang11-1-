import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import orgData from './mydata.json';
import cmsData from './cmsdata.json';
import staData from './ststusdata.json';
import appData from './appdata.json';
import intgsData from './intgStat.json'

declare module "*.json"{

}
interface Org{
  id:number;
  name:string;
}
interface cms1{
  id:number;
  name:string;
}
interface stat1{
  id:number;
  name:string;
}
interface app1{
  id:number;
  name:string;
}
interface intgs1{
  org_name:string;
  cms_name:string;
  stat_name:string;
  /*app_name:string;
  app_url:string;
  comm:string;*/
}
@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor(
    private httpClient:HttpClient
  ) { }

  ngOnInit(): void {
  }
 
  orglist: Org[]= orgData;
  Cms: cms1[]= cmsData;
  Stat: stat1[]= staData;
  App: app1[]= appData;
  Intgs: intgs1[]=intgsData;


 /* selectedOrg: string='';
  selectChangeHandlerOrg (event: any){
    this.selectedOrg =event.target.value;
  }
  selectedCms: string ='';
  selectChangeHandlerCms (event: any){
    this.selectedCms =event.target.value;
  }
  selectedSta: string= '';
  selectChangeHandlerSta (event: any){
    this.selectedSta =event.target.value;
  }
  addPost(){
    this.Intgs.push({
      org_name:this.selectedOrg,
      cms_name:this.selectedCms,
      stat_name:this.selectedSta
  })
}*/
  /*Org: any[] = [
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
  list=[
    {
      id:1, name:'Html5-SDK',
     
    },
    {
      id:2, name:'Windows App',
     
    },
    {
      id:3, name:'Android App',
     
    },
    {
      id:4, name:'Windows SDK',
    },
    {
      id:5, name:'Andriod SDK',
    },
    {
      id:6, name:'MaC App',
    },

   
  ]*/
  itemSelected(e:any){
    console.log(e);
  }
}
