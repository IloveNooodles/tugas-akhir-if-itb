export interface Company {
  id: string;
  name: string;
  cluster_name: string;
  created_at: string;
  updated_at: string;
}

export interface CompanyUser extends Company {
  username: string;
  email: string;
}
