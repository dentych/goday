using GDB.Model;
using Microsoft.EntityFrameworkCore;

namespace GDB.Database
{
    public class GoContext : DbContext
    {
        public DbSet<GoFile> GoFiles { get; set; }

        public GoContext(DbContextOptions options) : base(options) { }
        public GoContext() { }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            optionsBuilder.UseSqlServer(@"Server=(localdb)\mssqllocaldb;Database=Go;Trusted_Connection=True;");
        }
    }
}
