create table modules (
  moduleid varchar(40) primary key,
  modulename varchar(255) not null,
  status char(1) not null,
  path varchar(255),
  resourcekey varchar(255),
  icon varchar(255),
  sequence int not null,
  parent varchar(40),
  createdby varchar(40),
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp
);

create table users (
  userid varchar(40) primary key,
  username varchar(255) not null,
  email varchar(255) not null,
  displayname varchar(255) not null,
  status char(1) not null,
  gender char(1),
  phone varchar(20),
  title varchar(10),
  position varchar(40),
  imageurl varchar(500),
  createdby varchar(40),
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp,
  lastlogin timestamp
);

create table roles (
  roleid varchar(40) primary key,
  rolename varchar(255) not null,
  status char(1) not null,
  remark varchar(255),
  createdby varchar(40),
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp
);
create table userroles (
  userid varchar(40) not null,
  roleid varchar(40) not null,
  primary key (userid, roleid)
);
create table rolemodules (
  roleid varchar(40) not null,
  moduleid varchar(40) not null,
  permissions int not null,
  primary key (roleid, moduleid)
);

create table auditlog (
  id varchar(255) primary key,
  resourcetype varchar(255),
  userid varchar(255),
  ip varchar(255),
  action varchar(255),
  timestamp timestamp,
  status varchar(255),
  remark varchar(255)
);
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('dashboard','Dashboard','A','/dashboard','dashboard','assignments',1,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('admin','Admin','A','/admin','admin','contacts',2,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('setup','Setup','A','/setup','setup','settings',3,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('report','Report','A','/report','report','pie_chart',4,'');

insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('user','User Management','A','/admin/users','user','person',1,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('role','Role Management','A','/admin/roles','role','credit_card',2,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('audit_log','Audit Log','A','/admin/audit-logs','audit_log','zoom_in',4,'admin');

insert into roles (roleid, rolename, status, remark) values ('admin','Admin','A','Admin');
insert into roles (roleid, rolename, status, remark) values ('call_center','Call Center','A','Call Center');
insert into roles (roleid, rolename, status, remark) values ('it_support','IT Support','A','IT Support');
insert into roles (roleid, rolename, status, remark) values ('operator','Operator Group','A','Operator Group');

insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00001','gareth.bale','gareth.bale@gmail.com','Gareth Bale','https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00002','cristiano.ronaldo','cristiano.ronaldo@gmail.com','Cristiano Ronaldo','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00003','james.rodriguez','james.rodriguez@gmail.com','James Rodríguez','https://upload.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00004','zinedine.zidane','zinedine.zidane@gmail.com','Zinedine Zidane','https://upload.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00005','kaka','kaka@gmail.com','Kaká','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00006','luis.figo','luis.figo@gmail.com','Luís Figo','https://upload.wikimedia.org/wikipedia/commons/thumb/6/63/UEFA_TT_7209.jpg/440px-UEFA_TT_7209.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00007','ronaldo','ronaldo@gmail.com','Ronaldo','https://upload.wikimedia.org/wikipedia/commons/c/c8/Real_Valladolid-Valencia_CF%2C_2019-05-18_%2890%29_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00008','thibaut.courtois','thibaut.courtois@gmail.com','Thibaut Courtois','https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Courtois_2018_%28cropped%29.jpg/440px-Courtois_2018_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00009','luka.modric','luka.modric@gmail.com','Luka Modrić','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/ISL-HRV_%287%29.jpg/440px-ISL-HRV_%287%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00010','xabi.alonso','xabi.alonso@gmail.com','Xabi Alonso','https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg/440px-Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00011','karim.benzema','karim.benzema@gmail.com','Karim Benzema','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Karim_Benzema_2018.jpg/440px-Karim_Benzema_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00012','marc-andre.ter.stegen','marc-andre.ter.stegen@gmail.com','Marc-André ter Stegen','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e1/Marc-Andr%C3%A9_ter_Stegen.jpg/500px-Marc-Andr%C3%A9_ter_Stegen.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00013','sergino.dest','sergino.dest@gmail.com','Sergiño Dest','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Sergino_Dest.jpg/440px-Sergino_Dest.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00014','gerard.pique','gerard.pique@gmail.com','Gerard Piqué','https://upload.wikimedia.org/wikipedia/commons/4/4e/Gerard_Piqu%C3%A9_2018.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00015','ronald.araujo','ronald.araujo@gmail.com@gmail.com','Ronald Araújo','https://pbs.twimg.com/media/EtnqxaEU0AAc6A6.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00016','sergio.busquets','sergio.busquets@gmail.com@gmail.com','Sergio Busquets','https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/Sergio_Busquets_2018.jpg/440px-Sergio_Busquets_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00017','antoine.griezmann','antoine.griezmann@gmail.com@gmail.com','Antoine Griezmann','https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Antoine_Griezmann_in_2018_%28cropped%29.jpg/440px-Antoine_Griezmann_in_2018_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00018','miralem.pjanic','miralem.pjanic@gmail.com@gmail.com','Miralem Pjanić','https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg/440px-20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg','A','M','0987654321','Mrs','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00019','martin.braithwaite','martin.braithwaite@gmail.com@gmail.com','Martin Braithwaite','https://img.a.transfermarkt.technology/portrait/header/95732-1583334177.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00020','ousmane.dembele','ousmane.dembele@gmail.com@gmail.com','Ousmane Dembélé','https://upload.wikimedia.org/wikipedia/commons/7/77/Ousmane_Demb%C3%A9l%C3%A9_2018.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00021','riqui.puig','riqui.puig@gmail.com@gmail.com','Riqui Puig','https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Bar%C3%A7a_Napoli_12_%28cropped%29.jpg/440px-Bar%C3%A7a_Napoli_12_%28cropped%29.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00022','philip.coutinho','philip.coutinho@gmail.com@gmail.com','Philip Coutinho','https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Norberto_Murara_Neto_2019.jpg/440px-Norberto_Murara_Neto_2019.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00023','victor.lindelof','victor.lindelof@gmail.com@gmail.com','Victor Lindelöf','https://upload.wikimedia.org/wikipedia/commons/thumb/c/cc/CSKA-MU_2017_%286%29.jpg/440px-CSKA-MU_2017_%286%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00024','eric.bailly','eric.bailly@gmail.com@gmail.com','Eric Bailly','https://upload.wikimedia.org/wikipedia/commons/c/cf/Eric_Bailly_-_ManUtd.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00025','phil.jones','phil.jones@gmail.com@gmail.com','Phil Jones','https://upload.wikimedia.org/wikipedia/commons/f/fa/Phil_Jones_2018-06-28_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00026','harry.maguire','harry.maguire@gmail.com@gmail.com','Harry Maguire','https://upload.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00027','paul.pogba','paul.pogba@gmail.com@gmail.com','Paul Pogba','https://upload.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00028','edinson.cavani','edinson.cavani@gmail.com@gmail.com','Edinson Cavani','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Edinson_Cavani_2018.jpg/440px-Edinson_Cavani_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00029','juan.mata','juan.mata@gmail.com@gmail.com','Juan Mata','https://upload.wikimedia.org/wikipedia/commons/7/70/Ukr-Spain2015_%286%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00030','anthony.martial','anthony.martial@gmail.com@gmail.com','Anthony Martial','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Anthony_Martial_27_September_2017_cropped.jpg/440px-Anthony_Martial_27_September_2017_cropped.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00031','marcus.rashford','marcus.rashford@gmail.com@gmail.com','Marcus Rashford','https://upload.wikimedia.org/wikipedia/commons/5/5e/Press_Tren_CSKA_-_MU_%283%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00032','mason.greenwood','mason.greenwood@gmail.com@gmail.com','Mason Greenwood','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Mason_Greenwood.jpeg/440px-Mason_Greenwood.jpeg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00033','lee.grant','lee.grant@gmail.com@gmail.com','Lee Grant','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/LeeGrant09.jpg/400px-LeeGrant09.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00034','jesse.lingard','jesse.lingard@gmail.com@gmail.com','Jesse Lingard','https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/Jesse_Lingard_2018-06-13_1.jpg/440px-Jesse_Lingard_2018-06-13_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00035','keylor.navas','keylor.navas@gmail.com@gmail.com','Keylor Navas','https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Liver-RM_%288%29_%28cropped%29.jpg/440px-Liver-RM_%288%29_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00036','achraf.hakimi','achraf.hakimi@gmail.com@gmail.com','Achraf Hakimi','https://upload.wikimedia.org/wikipedia/commons/9/91/Iran-Morocco_by_soccer.ru_14_%28Achraf_Hakimi%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00037','presnel.kimpembe','presnel.kimpembe@gmail.com@gmail.com','Presnel Kimpembe','https://upload.wikimedia.org/wikipedia/commons/thumb/0/0e/Presnel_Kimpembe.jpg/400px-Presnel_Kimpembe.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00038','sergio.ramos','sergio.ramos@gmail.com@gmail.com','Sergio Ramos','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg/440px-FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00039','marquinhos','marquinhos@gmail.com@gmail.com','Marquinhos','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg/440px-Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00040','marco.verratti','marco.verratti@gmail.com@gmail.com','Marco Verratti','https://upload.wikimedia.org/wikipedia/commons/d/d0/Kiev-PSG_%289%29.jpg','A','M','0987654321','Mr','E');

insert into userroles(userid, roleid) values ('00001','admin');
insert into userroles(userid, roleid) values ('00003','admin');
insert into userroles(userid, roleid) values ('00004','admin');
insert into userroles(userid, roleid) values ('00005','it_support');
insert into userroles(userid, roleid) values ('00007','admin');
insert into userroles(userid, roleid) values ('00008','call_center');
insert into userroles(userid, roleid) values ('00009','it_support');
insert into userroles(userid, roleid) values ('00010','call_center');
insert into userroles(userid, roleid) values ('00011','it_support');
insert into userroles(userid, roleid) values ('00012','call_center');
insert into userroles(userid, roleid) values ('00012','it_support');

insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'setup', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'report', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'audit_log', 7);

insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'audit_log', 7);
/*
alter table userroles add foreign key (userid) references users (userid);
alter table userroles add foreign key (roleid) references roles (roleid);

alter table modules add foreign key (parent) references modules (moduleid);

alter table rolemodules add foreign key (roleid) references roles (roleid);
alter table rolemodules add foreign key (moduleid) references modules (moduleid);
*/